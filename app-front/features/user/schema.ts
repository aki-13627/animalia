import { z } from "zod";
import { postSchema, userBaseSchema } from "../post/schema";
import { petSchema } from "../pet/schema";

export type UserBase = z.infer<typeof userBaseSchema>;
export const userSchema = userBaseSchema.extend({
  email: z.string().email(),
  bio: z.string().min(0),
  posts: z.array(postSchema),
  pets: z.array(petSchema),
});

export type User = z.infer<typeof userSchema>;
