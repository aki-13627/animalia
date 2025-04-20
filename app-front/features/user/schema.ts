import { z } from 'zod';

export const userBaseSchema = z.object({
  id: z.string().uuid(),
  name: z.string(),
  iconImageUrl: z.string().url(),
});

export type UserBase = z.infer<typeof userBaseSchema>;
