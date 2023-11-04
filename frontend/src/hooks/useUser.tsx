

// Not Ready To Use
export const useUser = () => {
    const userId = localStorage.getItem("userId")
    const role = localStorage.getItem("role")
    return { userId, role }
}

export const useStudent = () => {
    const studentId = localStorage.getItem('studentId') ?? "";
    return studentId;
}

export const useProfessor = () => {
    const professorId = localStorage.getItem('professorId') ?? "";
    return professorId;
}