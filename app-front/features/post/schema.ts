import { z } from "zod";

// TODO: fix
export const userBaseSchema = z.object({
  id: z.string().uuid(),
  name: z.string(),
  iconImageUrl: z.string().url(),
});

export const postSchema = z.object({
  id: z.string().uuid(),
  caption: z.string().min(0),
  imageUrl: z.string().min(1),
  user: userBaseSchema,
  createdAt: z.string().datetime(),
});
