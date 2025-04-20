import { z } from "zod";
import { postSchema } from "../post/schema";
import { petSchema } from "../pet/schema";
import { userBaseSchema } from "../user/schema";

export const userSchema = userBaseSchema.extend({
  email: z.string().email(),
  bio: z.string().min(0),
  posts: z.array(postSchema),
  pets: z.array(petSchema),
});

export type User = z.infer<typeof userSchema>;

export const loginFormSchema = z.object({
  email: z
    .string()
    .email({ message: "有効なメールアドレスを入力してください" }),
  password: z.string().min(8, { message: "パスワードは8文字以上必要です" }),
});

export type LoginForm = z.infer<typeof loginFormSchema>;

export const loginResponseSchema = z.object({
  message: z.string(),
  user: userSchema,
  accessToken: z.string().min(1),
  idToken: z.string().min(1),
  refreshToken: z.string().min(1),
});

export type LoginResponse = z.infer<typeof loginResponseSchema>;

export const signUpResponseSchema = z.object({
  message: z.string(),
  user: userSchema,
});

export type SignUpResponse = z.infer<typeof signUpResponseSchema>;

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const signUpFormSchema = z.object({
  name: z.string(),
  email: z
    .string()
    .email({ message: "有効なメールアドレスを入力してください" }),
  password: z
    .string()
    .min(8, { message: "パスワードは8文字以上必要です" })
    .regex(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/, {
      message: "パスワードには大文字、小文字、数字を含める必要があります",
    }),
});

export type SignUpForm = z.infer<typeof signUpFormSchema>;

export const verifyEmailFormSchema = z.object({
  email: z
    .string()
    .email({ message: "有効なメールアドレスを入力してください" }),
  code: z.string().min(6, { message: "6桁のコードを入力してください" }),
});

export type VerifyEmailForm = z.infer<typeof verifyEmailFormSchema>;
