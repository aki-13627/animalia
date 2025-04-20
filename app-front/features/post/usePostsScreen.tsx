import { useAuth } from '@/providers/AuthContext';
import { fetchApi } from '@/utils/api';
import { useQuery } from '@tanstack/react-query';
import { z } from 'zod';
import { postSchema } from './schema';

const getPostsResponseSchema = z.object({
  posts: z.array(postSchema),
});

export const usePostsScreen = () => {
  const { token } = useAuth();
  const { data, isLoading, error, refetch } = useQuery({
    queryKey: ['posts'],
    queryFn: async () => {
      const result = await fetchApi({
        method: 'GET',
        path: 'posts',
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
