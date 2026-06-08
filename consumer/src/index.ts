import { connect, Subscription } from "nats";
import type { components } from "./types/events";
import { StringCodec } from "nats";

type EventMap = {
  "buff.listing.>": components["schemas"]["ListingEvent"];
  "dmarket.new_listing.>": components["schemas"]["NewListingEvent"];
  "dmarket.new_listing.*.FT-0.*": components["schemas"]["NewListingEvent"];
  "dmarket.orderbook.>": components["schemas"]["OrderbookEvent"];
  "dmarket.sales.>": components["schemas"]["SalesEvent"];
  "dmarket.orderbook.*.FT-0.*": components["schemas"]["OrderbookEvent"];
};

type EventSubject = keyof EventMap;

const NATS_URL = process.env.NATS_URL || "nats://localhost:4222";

const subjects: EventSubject[] = [
  "buff.listing.>",
  "dmarket.new_listing.>",
  "dmarket.new_listing.*.FT-0.*",
  "dmarket.orderbook.>",
  "dmarket.orderbook.*.FT-0.*",
  "dmarket.sales.>",
];

function handleListingEvent(event: EventMap["buff.listing.>"]) {
  const { item, source, event_timestamp } = event;
  // console.log(
  //   `[Listing] ${item.market_hash_name} | $${item.price} | ${item.exterior} | float: ${item.float_value ?? "N/A"}`,
  // );
  // console.log(`  source: ${source} | item_id: ${item.item_id}`);
  // if (item.stickers.length > 0) {
  //   console.log(`  stickers: ${item.stickers.map((s) => s.name).join(", ")}`);
  // }
}

function handleNewListingEvent(event: EventMap["dmarket.new_listing.>"]) {
  const { item, source, event_timestamp } = event;
  // console.log(
  //   `[New Listing] ${item.market_hash_name} | $${item.price} | ${item.exterior ?? "N/A"}`,
  // );
  // console.log(`  source: ${source} | offer_id: ${item.offer_id}`);
}

function handleOrderbookEvent(event: EventMap["dmarket.orderbook.>"]) {
  const { item, source, event_timestamp } = event;
  // console.log(
  //   `[Orderbook] ${item.market_hash_name} | $${item.price} | depth: ${item.depth}`,
  // );
  // console.log(`  source: ${source} | updated: ${item.updated_at}`);
}

function handleSalesEvent(event: EventMap["dmarket.sales.>"]) {
  const { item, source, event_timestamp } = event;
  // console.log(
  //   `[Sale] ${item.market_hash_name} | $${item.price} | ${item.sale_type}`,
  // );
  // console.log(`  source: ${source} | skin_id: ${item.skin_id}`);
}

function handleFT0NewListingEvent(
  event: EventMap["dmarket.new_listing.*.FT-0.*"],
) {
  const { item, source, event_timestamp } = event;
  console.log(
    `[New FT-0 Listing] ${item.market_hash_name} | $${item.price} | ${item.exterior ?? "N/A"}`,
  );
}

function handleOrderbookEventFT0(
  event: EventMap["dmarket.orderbook.*.FT-0.*"],
) {
  const { item, source, event_timestamp } = event;
  console.log(
    `[New FT-0 orderbook update] ${item.market_hash_name} | $${item.price}`,
  );
}

const handlers: Record<EventSubject, (event: unknown) => void> = {
  "buff.listing.>": (event) =>
    handleListingEvent(event as EventMap["buff.listing.>"]),
  "dmarket.new_listing.>": (event) =>
    handleNewListingEvent(event as EventMap["dmarket.new_listing.>"]),
  "dmarket.orderbook.>": (event) =>
    handleOrderbookEvent(event as EventMap["dmarket.orderbook.>"]),
  "dmarket.sales.>": (event) =>
    handleSalesEvent(event as EventMap["dmarket.sales.>"]),
  "dmarket.new_listing.*.FT-0.*": (event) =>
    handleFT0NewListingEvent(event as EventMap["dmarket.new_listing.*.FT-0.*"]),
  "dmarket.orderbook.*.FT-0.*": (event) =>
    handleOrderbookEventFT0(event as EventMap["dmarket.orderbook.>"]),
};

async function main() {
  console.log(`Connecting to NATS at ${NATS_URL}...`);
  const nc = await connect({ servers: NATS_URL });
  console.log("Connected to NATS");

  const sc = StringCodec();

  const subscriptions: Subscription[] = [];

  for (const subject of subjects) {
    const sub = nc.subscribe(subject, { queue: "consumer-group" });
    subscriptions.push(sub);

    (async () => {
      for await (const msg of sub) {
        try {
          const raw = sc.decode(msg.data);

          const event = JSON.parse(raw);
          handlers[subject](event);
        } catch (error) {
          console.error(`Error processing message on ${subject}:`, error);
        }
      }
    })();
  }

  console.log(`Subscribed to ${subjects.length} subjects:`);
  for (const subject of subjects) {
    console.log(`  - ${subject}`);
  }
  console.log("\nWaiting for events...\n");

  const stop = (signal: string) => {
    console.log(`\n${signal} received, shutting down...`);
    nc.close();
    process.exit(0);
  };

  process.on("SIGINT", () => stop("SIGINT"));
  process.on("SIGTERM", () => stop("SIGTERM"));
}

main().catch((error) => {
  console.error("Failed to start consumer:", error);
  process.exit(1);
});
