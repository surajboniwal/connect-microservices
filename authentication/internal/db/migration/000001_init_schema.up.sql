CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "fullname" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);
