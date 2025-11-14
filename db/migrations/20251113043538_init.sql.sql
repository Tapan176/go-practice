-- migrate:up

CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "firstName" VARCHAR(50) NOT NULL,
  "lastName" VARCHAR(50) NOT NULL,
  "email" VARCHAR(320) NOT NULL,
  "passwordHash" TEXT NOT NULL,
  "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "role" VARCHAR(20) NOT NULL DEFAULT 'user',
  "isVerified" boolean NOT NULL DEFAULT false,
  CONSTRAINT "cc_user_role" CHECK ("role" in ('admin', 'author', 'user'))
);

CREATE TABLE "articles" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(50) NOT NULL UNIQUE,
  "body" TEXT NOT NULL,
  "userId" INT NOT NULL,
  "categoryId" INT NOT NULL,
  "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "likes" INT NOT NULL DEFAULT 0,
  "dislikes" INT NOT NULL DEFAULT 0,
  FOREIGN KEY ("userId") REFERENCES users ("id"),
  FOREIGN KEY ("categoryId") REFERENCES categories ("id")
);

CREATE TABLE "comments" (
  "id" SERIAL PRIMARY KEY,
  "articleId" INT NOT NULL,
  "userId" INT NOT NULL,
  "comment" TEXT NOT NULL,
  "likes" INT NOT NULL DEFAULT 0,
  "dislikes" INT NOT NULL DEFAULT 0,
  "reply" JSONB NOT NULL DEFAULT '{}'::jsonb,
  "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updatedAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY ("articleId") REFERENCES articles ("id"),
  FOREIGN KEY ("userId") REFERENCES users ("id")
);

CREATE TABLE "likedAndDisliked" (
  "id" SERIAL PRIMARY KEY,
  "userId" INT NOT NULL, 
  "blogsLiked" INT[], 
  "blogsDisliked" INT[], 
  FOREIGN KEY ("userId") REFERENCES users ("id")
);

CREATE TABLE "likedBlog" (
  "userId" INT NOT NULL,
  "blogId" INT NOT NULL,
  FOREIGN KEY ("userId") REFERENCES USERS ("id"),
  FOREIGN KEY ("blogId") REFERENCES ARTICLES ("id")
);

CREATE TABLE "dislikedBlog" (
  "userId" INT NOT NULL,
  "blogId" INT NOT NULL,
  FOREIGN KEY ("userId") REFERENCES USERS ("id"),
  FOREIGN KEY ("blogId") REFERENCES ARTICLES ("id")
);

CREATE TABLE "authorRatings" (
  "userId" INT NOT NULL, 
  "authorId" INT NOT NULL, 
  "rating" INT NOT NULL, 
  FOREIGN KEY ("userId") REFERENCES "users" ("id"), 
  FOREIGN KEY ("authorId") REFERENCES "users" ("id"), 
  CONSTRAINT "cc_user_rating" CHECK ("rating" > 0 AND "rating" < 6)
);

-- migrate:down

