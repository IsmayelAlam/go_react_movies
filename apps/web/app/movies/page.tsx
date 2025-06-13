'use client';

import { axiosCall } from '@/lib/axios';
import { useQuery } from '@tanstack/react-query';

export default function Movies() {
  const { data: movies, isLoading: isLoadingMovies } = useQuery({
    queryKey: ['movies'],
    queryFn: () =>
      axiosCall({
        method: 'get',
        endpoint: 'movies',
      }),
  });

  console.log('Movies:', movies, isLoadingMovies);

  return <div>Movies</div>;
}
