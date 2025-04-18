import { useCallback, useState } from "react";
import { useStudentsGetAll } from "../../query/useGetAll";
import { useNavigate } from "react-router-dom";
import { LIMIT_DEFAULT, PAGE_DEFAULT } from "./constants";

const StudentListScreen = () => {
    const [page, setPage] = useState(PAGE_DEFAULT);
    const { data, isLoading } = useStudentsGetAll({
        page,
        limit: LIMIT_DEFAULT
    });  

    const navigate = useNavigate();

    const goTo = useCallback(() => {
        navigate("/form");
    }, [navigate]);

    const onNext = useCallback(() => {
        setPage((prev) => prev + 1);
    }, []);

    const onPrev = useCallback(() => {
        setPage((prev) => Math.max(prev - 1, 1));
    }, []);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (!data) {
        return <div>No data available</div>;
    }

    return (
        <div>
            <h1>Students List</h1>
            List: <button onClick={goTo}>+</button>            
            {data && data.students.map((student) => (<p key={student.id}>- {student.name}</p>))}     
            <p>Total Students: {data ? data.total : 0}</p>        
            
            <p>
                <button onClick={onPrev} disabled={page <= 1}>Previous</button>
                <span> {page} </span>
                <button onClick={onNext} disabled={page >= data.totalPages}>Next</button>
            </p>                  
        </div>
    );
}

export default StudentListScreen;