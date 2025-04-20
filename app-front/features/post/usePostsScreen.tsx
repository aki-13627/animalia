import { useAuth } from "@/providers/AuthContext";
import { fetchApi } from "@/utils/api";
import { useQuery } from "@tanstack/react-query";
import { z } from "zod";

const userBaseSchema = z.object({
  id: z.string().uuid(),
  name: z.string(),
  iconImageUrl: z.string().url(),
});

const postSchema = z.object({
  id: z.string().uuid(),
  caption: z.string().min(0),
  imageUrl: z.string().min(1),
  user: userBaseSchema,
  createdAt: z.string().datetime(),
});

const getPostsResponseSchema = z.object({
  posts: z.array(postSchema),
});

export const usePostsScreen = () => {
  const { token } = useAuth();
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ["posts"],
    queryFn: async () => {
      const result = await fetchApi({
        path: "posts",
        schema: getPostsResponseSchema,
        options: {},
        token,
      });
      return result.posts;
    },
  });

  return {
    data: data,
    isLoading,
    error,
    refetch,
  };
};
