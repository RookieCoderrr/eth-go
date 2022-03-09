CREATE TABLE "public"."balances" (
  "address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "eth_balance" numeric(64) NOT NULL,
  "block_number" int4 NOT NULL
);
CREATE INDEX "balances_address_9c33318b_like" ON "public"."balances" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);
CREATE INDEX "idx_balances_address" ON "public"."balances" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_balances_block_number" ON "public"."balances" USING btree (
  "block_number"
);
ALTER TABLE "public"."balances" ADD CONSTRAINT "balances_blockNumber_ce829c90_check" CHECK (("block_number" >= 0));
ALTER TABLE "public"."balances" ADD CONSTRAINT "balances_pkey" PRIMARY KEY ("address");



CREATE TABLE "public"."blocks" (
  "number" int8 NOT NULL,
  "hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "difficulty" numeric(64) NOT NULL,
  "extra_data" text COLLATE "pg_catalog"."default" NOT NULL,
  "gas_limit" int4 NOT NULL,
  "gas_used" int4 NOT NULL,
  "logs_bloom" text COLLATE "pg_catalog"."default" NOT NULL,
  "miner" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "mix_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "nonce" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "parent_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "receipts_root" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "sha3_uncles" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "size" int4 NOT NULL,
  "state_root" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "timestamp" int4 NOT NULL,
  "total_difficulty" numeric(64) NOT NULL,
  "transactions_root" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "uncles" text COLLATE "pg_catalog"."default" NOT NULL,
  "transaction_count" int4 NOT NULL
);
CREATE INDEX "blocks_hash_b84f947b" ON "public"."blocks" USING btree (
  "hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "blocks_hash_b84f947b_like" ON "public"."blocks" USING btree (
  "hash" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);
CREATE INDEX "blocks_miner_4c97a59a" ON "public"."blocks" USING btree (
  "miner" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "blocks_miner_4c97a59a_like" ON "public"."blocks" USING btree (
  "miner" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);
CREATE INDEX "blocks_mix_hash_f2ebab47" ON "public"."blocks" USING btree (
  "mix_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "blocks_mix_hash_f2ebab47_like" ON "public"."blocks" USING btree (
  "mix_hash" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);

ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_number_check" CHECK ((number >= 0));
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_transactionCount_check" CHECK (("transaction_count" >= 0));
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_size_check" CHECK ((size >= 0));
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_gasLimit_check" CHECK (("gas_limit" >= 0));
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_gasUsed_check" CHECK (("gas_used" >= 0));
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_timestamp_check" CHECK (("timestamp" >= 0));
ALTER TABLE "public"."blocks" ADD CONSTRAINT "blocks_pkey" PRIMARY KEY ("number");


CREATE TABLE "public"."contracts" (
  "address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default",
  "bytecode" text COLLATE "pg_catalog"."default" NOT NULL,
  "abi" text COLLATE "pg_catalog"."default",
  "runs" int4 NOT NULL,
  "bytecode_hash" varchar(128) COLLATE "pg_catalog"."default",
  "source" text COLLATE "pg_catalog"."default",
  "compiler" varchar(128) COLLATE "pg_catalog"."default",
  "library" varchar(1024) COLLATE "pg_catalog"."default",
  "constructor_args" text COLLATE "pg_catalog"."default",
  "block_number" int4 NOT NULL,
  "timestamp" int4 NOT NULL,
  "creator" varchar(64) COLLATE "pg_catalog"."default",
  "transaction_hash" varchar(128) COLLATE "pg_catalog"."default"
)
;
CREATE INDEX "contracts_address_7cd54651_like" ON "public"."contracts" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);
CREATE INDEX "idx_contracts_address" ON "public"."contracts" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_contracts_block_number" ON "public"."contracts" USING btree (
  "block_number"
);
ALTER TABLE "public"."contracts" ADD CONSTRAINT "contracts_runs_check" CHECK ((runs >= 0));
ALTER TABLE "public"."contracts" ADD CONSTRAINT "contracts_timestamp_check" CHECK (("timestamp" >= 0));
ALTER TABLE "public"."contracts" ADD CONSTRAINT "contracts_blockNumber_check" CHECK (("block_number" >= 0));
ALTER TABLE "public"."contracts" ADD CONSTRAINT "contracts_pkey" PRIMARY KEY ("address");


CREATE SEQUENCE logs_id_seq;
CREATE TABLE "public"."tx_logs" (
  "id" int8 NOT NULL DEFAULT nextval('logs_id_seq'::regclass),
  "address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "block_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "event_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "block_number" int4 NOT NULL,
  "data" text COLLATE "pg_catalog"."default" NOT NULL,
  "log_index" int4 NOT NULL,
  "removed" bool NOT NULL,
  "topics" text COLLATE "pg_catalog"."default" NOT NULL,
  "transaction_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "transaction_index" int4 NOT NULL,
  "args" jsonb,
  "topic0" varchar(128) COLLATE "pg_catalog"."default" NOT NULL
);
ALTER TABLE "public"."tx_logs" ADD CONSTRAINT "logs_blockNumber_check" CHECK (("block_number" >= 0));
ALTER TABLE "public"."tx_logs" ADD CONSTRAINT "logs_logIndex_check" CHECK (("log_index" >= 0));
ALTER TABLE "public"."tx_logs" ADD CONSTRAINT "logs_transactionIndex_check" CHECK (("transaction_index" >= 0));
ALTER TABLE "public"."tx_logs" ADD CONSTRAINT "logs_pkey" PRIMARY KEY ("id");
CREATE INDEX "idx_logs_address" ON "public"."tx_logs" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_logs_topic0" ON "public"."tx_logs" USING btree (
  "topic0" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
-- 不锁表创建索引
CREATE INDEX CONCURRENTLY "idx_logs_block_number" ON "public"."tx_logs" USING btree (
  "block_number"
);
-- 不锁表创建联合索引
CREATE INDEX CONCURRENTLY "idx_logs_log_index_block_number" ON "public"."tx_logs" USING btree (
  "block_number", "log_index"
);


CREATE SEQUENCE token_transfers_id_seq;
CREATE TABLE "public"."token_transfers" (
  "id" int8 NOT NULL DEFAULT nextval('token_transfers_id_seq'::regclass),
  "token_address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "from_address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "to_address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "value" numeric(128) NOT NULL,
  "transaction_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "log_index" int4 NOT NULL,
  "block_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "block_number" int4 NOT NULL,
  "event_name" varchar(64) COLLATE "pg_catalog"."default",
  "topic" varchar(128) COLLATE "pg_catalog"."default",
  "timestamp" int4,
  "transfer_type" int4 NOT NULL
);
ALTER TABLE "public"."token_transfers" ADD CONSTRAINT "token_transfers_logIndex_check" CHECK (("log_index" >= 0));
ALTER TABLE "public"."token_transfers" ADD CONSTRAINT "token_transfers_blockNumber_check" CHECK (("block_number" >= 0));
ALTER TABLE "public"."token_transfers" ADD CONSTRAINT "token_transfers_timestamp_check" CHECK (("timestamp" >= 0));
ALTER TABLE "public"."token_transfers" ADD CONSTRAINT "token_transfers_pkey" PRIMARY KEY ("id");
-- 不锁表创建索引
CREATE INDEX CONCURRENTLY "idx_token_transfers_block_number" ON "public"."token_transfers" USING btree (
  "block_number"
);


CREATE TABLE "public"."tokens" (
  "address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "symbol" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "decimals" int2 NOT NULL,
  "is_erc20" bool NOT NULL,
  "is_erc721" bool NOT NULL,
  "total_supply" numeric(64,10) NOT NULL,
  "block_timestamp" int4 NOT NULL,
  "block_number" int4 NOT NULL,
  "block_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL
);
CREATE INDEX "idx_tokens_address" ON "public"."tokens" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "tokens_address_be80e124_like" ON "public"."tokens" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);
ALTER TABLE "public"."tokens" ADD CONSTRAINT "tokens_blockTimestamp_check" CHECK (("block_timestamp" >= 0));
ALTER TABLE "public"."tokens" ADD CONSTRAINT "tokens_blockNumber_check" CHECK (("block_number" >= 0));
ALTER TABLE "public"."tokens" ADD CONSTRAINT "tokens_pkey" PRIMARY KEY ("address");



CREATE TABLE "public"."transcations" (
  "transaction_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "block_hash" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "block_number" int8 NOT NULL,
  "transaction_index" int4 NOT NULL,
  "tx_type" text COLLATE "pg_catalog"."default" NOT NULL,
  "from_address" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "to_address" varchar(64) COLLATE "pg_catalog"."default",
  "gas" int8 NOT NULL,
  "cumulative_gas_used" int4 NOT NULL,
  "gas_used" int4 NOT NULL,
  "gas_price" numeric(64) NOT NULL,
  "nonce" int4 NOT NULL,
  "value" numeric(64) NOT NULL,
  "r" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "s" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "v" int4 NOT NULL,
  "status" int2 NOT NULL,
  "timestamp" int4 NOT NULL,
  "logs_bloom" text COLLATE "pg_catalog"."default" NOT NULL,
  "contract_address" varchar(128) COLLATE "pg_catalog"."default",
  "log_count" int4 NOT NULL,
  "input" text COLLATE "pg_catalog"."default"
)
;
CREATE INDEX "idx_transcations_transaction_hash" ON "public"."transcations" USING btree (
  "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
-- 不锁表创建索引
CREATE INDEX CONCURRENTLY "idx_transcations_block_number" ON "public"."transcations" USING btree (
  "block_number"
);
CREATE INDEX "transcations_hash_a916c837_like" ON "public"."transcations" USING btree (
  "transaction_hash" COLLATE "pg_catalog"."default" "pg_catalog"."varchar_pattern_ops" ASC NULLS LAST
);
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_v_check" CHECK ((v >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_timestamp_check" CHECK (("timestamp" >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_blockNumber_check" CHECK (("block_number" >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_gas_check" CHECK ((gas >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_cumulativeGasUsed_check" CHECK (("cumulative_gas_used" >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_nonce_check" CHECK ((nonce >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_transactionIndex_check" CHECK (("transaction_index" >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_gasUsed_check" CHECK (("gas_used" >= 0));
ALTER TABLE "public"."transcations" ADD CONSTRAINT "transcations_pkey" PRIMARY KEY ("transaction_hash","block_number");




-- 如果有需要再执行以下sql，及得修改用户名tbaseadmin为自己的用户名，一般不需要执行
-- ALTER TABLE "public"."balances" OWNER TO "tbaseadmin";
-- ALTER TABLE "public"."contracts" OWNER TO "tbaseadmin";
-- ALTER TABLE "public"."blocks" OWNER TO "tbaseadmin";
-- ALTER TABLE "public"."tx_logs" OWNER TO "tbaseadmin";
-- ALTER TABLE "public"."token_transfers" OWNER TO "tbaseadmin";
-- ALTER TABLE "public"."tokens" OWNER TO "tbaseadmin";
-- ALTER TABLE "public"."transcations" OWNER TO "tbaseadmin";