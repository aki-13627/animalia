import { z } from "zod";
import { userBaseSchema } from "../user/schema";

export const postSchema = z.object({
  id: z.string().uuid(),
  caption: z.string().min(0),
  imageUrl: z.string().min(1),
  user: userBaseSchema,
  createdAt: z.string().datetime(),
});
