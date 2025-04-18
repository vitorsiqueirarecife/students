import { useQuery } from '@tanstack/react-query';
import axiosApi from '../../../../shared/api';
import { queryKeyStudentsGetAll } from './keys';
import { GetAllStudentsOptions, GetAllStudentsResponse } from '../../entities/types';




const queryFn = async ({ page, limit }: GetAllStudentsOptions): Promise<GetAllStudentsResponse> => {
    const response = await axiosApi.get<GetAllStudentsResponse>(`/students?page=${page}&limit=${limit}`);
    return response.data;
}

export const useStudentsGetAll = ({ page, limit }: GetAllStudentsOptions = { page: 1, limit: 10 }) => {
  return useQuery<GetAllStudentsResponse>({
    queryKey: queryKeyStudentsGetAll,
    queryFn: () => queryFn({ page, limit }),
  });
};
