// src/hooks/useApi.ts
import { useQuery, useMutation } from '@tanstack/react-query'
import api from '../api/axios'

export const useGet = (endpoint: string) => {
  return useQuery({
    queryKey: [endpoint],
    queryFn: () => api.get(endpoint).then(res => res.data)
  })
}

export const usePost = (endpoint: string) => {
  return useMutation({
    mutationFn: (data: any) => api.post(endpoint, data).then(res => res.data)
  })
}