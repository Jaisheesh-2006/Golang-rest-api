import React, { useEffect, useState } from "react";
import { Student } from "../types";
import StudentCard from "./StudentCard";
import EditStudentModal from "./EditStudentModal";
import { fetchStudents, deleteStudent } from "../services/api";
import "./StudentList.css";

interface StudentListProps {
  onRefresh?: () => void;
}

const StudentList: React.FC<StudentListProps> = ({ onRefresh }) => {
  const [students, setStudents] = useState<Student[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [editingStudent, setEditingStudent] = useState<Student | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const loadStudents = async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await fetchStudents();
      setStudents(data);
    } catch (err) {
      setError("Failed to fetch students");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadStudents();
  }, []);

  const handleEdit = (student: Student) => {
    setEditingStudent(student);
    setIsModalOpen(true);
  };

  const handleStudentUpdated = (updatedStudent: Student) => {
    setStudents((prevStudents) =>
      prevStudents.map((s) => (s.id === updatedStudent.id ? updatedStudent : s))
    );
  };

  const handleDelete = async (id: number) => {
    const success = await deleteStudent(id);
    if (success) {
      setStudents((prevStudents) => prevStudents.filter((s) => s.id !== id));
    } else {
      alert("Failed to delete student");
    }
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
    setEditingStudent(null);
  };

  if (loading) {
    return <div className="loading">Loading students...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (students.length === 0) {
    return (
      <div className="empty-state">
        No students yet. Add one to get started!
      </div>
    );
  }

  return (
    <div className="student-list">
      <h2>Students ({students.length})</h2>
      <div className="cards-container">
        {students.map((student) => (
          <StudentCard
            key={student.id}
            student={student}
            onEdit={handleEdit}
            onDelete={handleDelete}
          />
        ))}
      </div>
      {editingStudent && (
        <EditStudentModal
          student={editingStudent}
          isOpen={isModalOpen}
          onClose={handleCloseModal}
          onStudentUpdated={handleStudentUpdated}
        />
      )}
    </div>
  );
};

export default StudentList;
