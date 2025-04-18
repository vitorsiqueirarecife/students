import { useForm } from 'react-hook-form';
import { Student } from '../../entities/types';
import { useCallback } from 'react';

const StudentFormScreen = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Student>();

  const onSubmit = useCallback((data: Student) => {
    console.log('Student Data:', data);
  }, []);

  return (
    <div>
      <h2>Student Form</h2>
      <form onSubmit={handleSubmit(onSubmit)}>   
        <div style={{ marginBottom: '10px' }}>
          <label>Name</label>
          <input {...register('name', { required: 'Name is required' })} />
          {errors.name && <p>{errors.name.message}</p>}
        </div>

        <div style={{ marginBottom: '10px' }}>
          <label>Grade</label>
          <input {...register('grade', { required: 'Grade is required' })} />
          {errors.grade && <p>{errors.grade.message}</p>}
        </div>

        <button type="submit">Salvar</button>
      </form>
    </div>
  );
};

export default StudentFormScreen;
