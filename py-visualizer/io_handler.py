import re
from abc import ABC, abstractmethod
from pathlib import Path

import pandas as pd


class IOHandler(ABC):
    @abstractmethod
    def get_providers_status(self) -> dict[int, dict[int, bool]]:
        """
        Returns a dictionary containing the status of providers and their associated
        block numbers. The status is represented as a boolean value, where True
        indicates a failure and False indicates success.
        """
        raise NotImplementedError

    @abstractmethod
    def get_price_feeds(self) -> pd.DataFrame:
        """
        Returns price feeds data with block numbers as the index and each provider
        as a separate column in the DataFrame.
        """
        raise NotImplementedError


class LogfileIOHandler(IOHandler):
    def __init__(self, logpath: Path) -> None:
        super().__init__()
        self.logpath = logpath

    def get_log_paths(self) -> tuple[Path, dict[int, Path]]:
        """
        Get the paths of log files for admin daemon and fast update providers.

        Args:
            filepath (Path): The directory path where the log files are located.

        Returns:
            tuple[Path, dict[int, Path]]: A tuple containing the path of the admin
            daemon log file and a dictionary mapping provider IDs to their
            corresponding log files.
        Raises:
            FileNotFoundError: If no fast-update-provider logs or admin-daemon logs
                are found.
        """
        admin_daemon_log: Path = Path()
        fast_update_provider_logs: dict[int, Path] = {}

        for file in self.logpath.iterdir():
            if "fast-update-provider" in file.stem:
                provider_id = int(re.findall(r"\d+", file.stem)[0])
                fast_update_provider_logs[provider_id] = file
            if "admin-daemon" in file.stem:
                admin_daemon_log = file

        if len(fast_update_provider_logs) == 0:
            raise FileNotFoundError("No fast-update-provider logs found")
        if not admin_daemon_log:
            raise FileNotFoundError("No admin-daemon logs found")

        return admin_daemon_log, fast_update_provider_logs

    def get_providers_status(self) -> dict[int, dict[int, bool]]:
        """
        Retrieve the status of providers based on the contents of the files
        specified by the filepaths.

        Args:
            filepath (Dict[int, Path]): A dictionary mapping provider IDs to file paths.

        Returns:
            Dict[int, Dict[int, bool]]: A dictionary containing the status of providers
            and their associated block numbers.The status is represented as a boolean
            value, where True indicates a failure and False indicates success.
        """
        _, filepath = self.get_log_paths()

        status = {}

        for provider_id, provider_path in filepath.items():
            status[provider_id] = {}
            with open(provider_path) as file:
                for line in file.readlines():
                    if "Block" not in line:
                        continue
                    nums = re.findall(r"\d+", line)
                    block_number = int(nums[8])
                    if "failed" in line or "Failed" in line or "Error" in line:
                        status[provider_id][block_number] = True
                    else:
                        status[provider_id][block_number] = False
        return status

    def get_providers_price_comparison(
        self,
        filepaths: dict[int, Path],
    ) -> dict[int, pd.DataFrame]:
        """
        Retrieves price comparison data for different providers.

        Args:
            filepaths (dict[int, Path]): A dictionary mapping provider IDs to
                file paths.

        Returns:
            dict[int, pd.DataFrame]: A dictionary mapping provider IDs to pandas
                DataFrames containing price comparison data.
        """
        dfs = {}
        for provider_id, filepath in filepaths.items():
            block_numbers, feeds = [], []
            with open(filepath) as file:
                for line in file.readlines():
                    if "chainPrices" not in line:
                        continue
                    nums = re.findall(r"\d+", line)
                    block_numbers.append(int(nums[8]))
                    feeds.append([int(nums[9]), int(nums[10])])

            df = pd.DataFrame(feeds, columns=["fast_update_price", "actual_price"])
            df["block_number"] = block_numbers
            df.set_index("block_number", drop=True, inplace=True)
            dfs[provider_id] = df
        return dfs

    def get_price_feeds(self) -> pd.DataFrame:
        """
        Parses a logfile and extracts price feeds data.

        Args:
            filepath (Path): The path to the logfile.
            skip_lines (int, optional): The number of lines to skip at
                the beginning of the file. Defaults to 0.

        Returns:
            pd.DataFrame: A DataFrame containing the parsed price feeds data.
        """
        filepath, _ = self.get_log_paths()

        block_numbers, fast_update_feeds, actual_feeds = [], [], []
        with open(filepath) as file:
            for line in file.readlines():
                if "Block" not in line:
                    continue
                nums = re.findall(r"(?:\d+\.)?\d+", line)
                num_feeds = int(nums[6])
                block_numbers.append(int(nums[7]))
                fast_update_feeds.append(
                    [float(num) for num in nums[8 : 8 + num_feeds]]
                )
                actual_feeds.append(
                    [float(num) for num in nums[8 + num_feeds : 8 + 8 + num_feeds]]
                )

        # Convert to DataFrame
        df = pd.DataFrame(fast_update_feeds)
        df["block_number"] = block_numbers
        df.set_index("block_number", drop=True, inplace=True)
        df = df.add_prefix("FastUpdateFeed_")

        df2 = pd.DataFrame(actual_feeds)
        df2["block_number"] = block_numbers
        df2.set_index("block_number", drop=True, inplace=True)
        df2 = df2.add_prefix("ActualFeed_")

        df = pd.concat([df, df2], axis=1)
        return df
