import { randomInt } from "crypto";
import { Feed } from "../protocol/voting-types";
import { getLogger } from "../utils/logger";

export class PriceFeedProvider {
    constructor(private readonly numFeeds: number) {
        this.numFeeds = numFeeds;
    }

    getFeed(): string {
        let feeds: string = "";
        let feed = 0;
        for (let i = 0; i < this.numFeeds; i++) {
            if (i % 2 == 0) {
                feed = 0;
            }
            const n = randomInt(3);
            if (i % 2 == 0) {
                if (n == 1) {
                    feed += 1;
                }
                if (n == 2) {
                    feed += 3;
                }
            } else {
                if (n == 1) {
                    feed += 4;
                }
                if (n == 2) {
                    feed += 12;
                }
                feeds = feeds + feed.toString(16);
            }
        }
        feeds = feeds + "0".repeat(500 - feeds.length);

        return feeds;
    }
}
