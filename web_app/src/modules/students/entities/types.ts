export type Student = {
    id: string;
    name: string;
    grade: string;
};

export type GetAllStudentsResponse = {
    students: Student[];
    total: number;
    totalPages: number;
};

export type GetAllStudentsOptions = {
    page: number;
    limit: number;
};