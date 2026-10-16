import json
from abc import ABC, abstractmethod
from collections.abc import Iterator
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

    def get_entries(self) -> Iterator[dict]:
        """
        Read the Go client log entries. The client must log in the json format.

        Yields:
            dict: One parsed log entry per line. Lines that are not JSON are skipped.
        """
        with open(self.get_log_path()) as file:
            for line in file:
                try:
                    yield json.loads(line)
                except json.JSONDecodeError:
                    continue

    def get_provider_status(self) -> dict[int, bool]:
        """
        Retrieve the status of provider based on the contents of the file
        specified by the filepath.

        Returns:
            Dict[int, bool]: A dictionary containing the status of the provider
            and their associated block numbers. The status is represented as a boolean
            value, where True indicates a failure and False indicates success.
        """
        status = {}
        for entry in self.get_entries():
            if entry.get("msg") == "Scheduling update":
                status[str(entry["block"])] = True
            elif entry.get("msg") == "Update included":
                status[str(entry["block"])] = False
        return status

    def get_price_feeds(self) -> pd.DataFrame:
        """
        Parses a logfile and extracts price feeds data.

        Returns:
            pd.DataFrame: A DataFrame containing the parsed price feeds data.
        """
        block_numbers, fast_update_feeds, actual_feeds = [], [], []
        for entry in self.get_entries():
            # on-chain and provider values in a block
            if entry.get("msg") == "Feed values":
                block_numbers.append(entry["block"])
                fast_update_feeds.append(entry["chain"])
                actual_feeds.append(entry["provider"])

            # on-chain values after an update, without provider values
            elif entry.get("msg") == "Feed values after update":
                block_numbers.append(entry["block"])
                fast_update_feeds.append(entry["chain"])
                actual_feeds.append([None for _ in entry["chain"]])

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
        for entry in self.get_entries():
            if entry.get("msg") == "Feed ids fetched":
                return {
                    idx: f"{idx} - {name or 'delisted'}"
                    for idx, name in enumerate(entry["feeds"])
                }
        return {}
