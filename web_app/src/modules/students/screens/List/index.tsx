import { useStudentsGetAll } from "../../query/useGetAll";

const List = () => {

    const { data } = useStudentsGetAll();    

    return (
        <div>
            <h1>Estudiantes</h1>

            {data && data.map((student) => (student.name))}
        </div>
    );
}

export default List;