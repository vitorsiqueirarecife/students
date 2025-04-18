import { useCallback } from "react";
import { useStudentsGetAll } from "../../query/useGetAll";
import { useNavigate } from "react-router-dom";

const StudentListScreen = () => {
    const { data } = useStudentsGetAll();  
    const navigate = useNavigate();

    const goTo = useCallback(() => {
        navigate("/form");
    }, [navigate]);

    return (
        <div>
            <h1>Estudiantes</h1>
            List: <button onClick={goTo}>+</button>            
            {data && data.map((student) => (<p>{student.name}</p>))}            
        </div>
    );
}

export default StudentListScreen;