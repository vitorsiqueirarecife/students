import { useQuery } from '@tanstack/react-query';
import axiosApi from '../../../../shared/api';
import { queryKeyStudentsGetAll } from './keys';
import { Student } from '../../entities/types';


const queryFn = async (): Promise<Student[]> => {
    const response = await axiosApi.get<Student[]>('/students');
    return response.data;
}

export const useStudentsGetAll = () => {
  return useQuery<Student[]>({
    queryKey: queryKeyStudentsGetAll,
    queryFn: queryFn
  });
};
