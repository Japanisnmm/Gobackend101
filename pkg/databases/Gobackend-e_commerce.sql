CREATE TABLE "users" (
  "id" varchar PRIMARY,
  "username" varchar UNIQUE,
  "password" varchar,
  "email" varchar UNIQUE,
  "role_id" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "oauth" (
  "id" varchar PRIMARY,
  "user_id" varchar,
  "access_token" varchar,
  "refresh_token" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "roles" (
  "id" int PRIMARY KEY,
  "title" varchar
);

CREATE TABLE "products" (
  "id" varchar PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "price" float,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "images" (
  "id" varchar,
  "filename" varchar,
  "url" varchar,
  "product_id" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "products_categories" (
  "id" varchar PRIMARY KEY,
  "products_id" varchar,
  "categories_id" int
);

CREATE TABLE "categories" (
  "id" int PRIMARY KEY,
  "title" varchar UNIQUE
);

CREATE TABLE "orders" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "contact" varchar,
  "address" varchar,
  "transfer_slip" json,
  "status" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "products_order" (
  "id" varchar PRIMARY KEY,
  "order_id" varchar,
  "qty" int,
  "product" json
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "oauth" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("categories_id") REFERENCES "categories" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "products_order" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
