import re
from abc import ABC, abstractmethod
from pathlib import Path

import pandas as pd


class IOHandler(ABC):
    @abstractmethod
    def get_provider_status(self) -> dict[int, bool]:
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

    def get_log_path(self) -> Path:
        """
        Get the path of Go client log file.

        Returns:
            Path: The path of the Go client log file.
        Raises:
            FileNotFoundError: If no fast_updates_client logs are found.
        """
        for file in self.logpath.iterdir():
            if "fast_updates_client" in file.stem:
                log_path = file
                break

        if not log_path:
            raise FileNotFoundError("No fast_updates_client logs found")

        return log_path

    def get_provider_status(self) -> dict[int, bool]:
        """
        Retrieve the status of provider based on the contents of the file
        specified by the filepath.

        Returns:
            Dict[int, bool]: A dictionary containing the status of the provider
            and their associated block numbers. The status is represented as a boolean
            value, where True indicates a failure and False indicates success.
        """
        provider_path = self.get_log_path()

        status = {}

        with open(provider_path) as file:
            block_number = None
            for line in file.readlines():
                if "scheduling update" in line:
                    block_number = re.findall(r"\d+", line)[7]
                    status[block_number] = True
                if "successful update for block" in line:
                    block_number = re.findall(r"\d+", line)[7]
                    status[block_number] = False
        return status

    def get_price_feeds(self) -> pd.DataFrame:
        """
        Parses a logfile and extracts price feeds data.

        Returns:
            pd.DataFrame: A DataFrame containing the parsed price feeds data.
        """
        filepath = self.get_log_path()

        block_numbers, fast_update_feeds, actual_feeds = [], [], []
        with open(filepath) as file:
            for line in file.readlines():
                # new block number + fast update feeds
                if "chain feeds" in line and "after update" not in line:
                    nums = re.findall(r"(?:\d+\.)?\d*(?:e?-?\d+)", line)
                    block_numbers.append(int(nums[5]))
                    fast_update_feeds.append([float(num) for num in nums[6:]])

                # get actual feeds for new block
                elif "provider feeds values" in line:
                    nums = re.findall(r"(?:\d+\.)?\d*(?:e?-?\d+)", line)
                    num_feeds = len(fast_update_feeds[0])
                    actual_feeds.append([float(num) for num in nums[5 : num_feeds + 5]])

                # write down chain values after update
                elif "after update" in line:
                    nums = re.findall(r"(?:\d+\.)?\d*(?:e?-?\d+)", line)
                    block_numbers.append(int(nums[5]))
                    fast_update_feeds.append([float(num) for num in nums[6:]])
                    actual_feeds.append([None for _ in nums[6:]])

        # Convert to DataFrame
        df = pd.DataFrame(fast_update_feeds)
        df["block_number"] = block_numbers
        df.set_index("block_number", drop=True, inplace=True)
        df = df.add_prefix("FastUpdateFeed_")
        # Sort by block number
        df.sort_index(inplace=True)

        df2 = pd.DataFrame(actual_feeds)
        df2["block_number"] = block_numbers
        df2.set_index("block_number", drop=True, inplace=True)
        df2 = df2.add_prefix("ActualFeed_")
        # Sort by block number
        df2.sort_index(inplace=True)

        df = pd.concat([df, df2], axis=1)
        return df

    def get_feed_names(self) -> dict[int, str]:
        """
        Get the names of the feeds from the log files.

        Returns:
            dict[int, str]: A dictionary mapping feed indices to their names.
        """
        filepath = self.get_log_path()

        feed_names = {}
        with open(filepath) as file:
            for line in file.readlines():
                if "Fetched feed ids" in line:
                    # the regex pattern is AAA(A)/AAA(A)
                    feed_names = re.findall(r"[A-Z]{3,4}\/[A-Z]{3,4}", line)
                    feed_names = {
                        idx: f"{idx} - " + name for idx, name in enumerate(feed_names)
                    }
                    break
        return feed_names
