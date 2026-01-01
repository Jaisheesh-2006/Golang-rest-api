import React, { useState } from "react";
import { Student } from "../types";
import "./StudentCard.css";

interface StudentCardProps {
  student: Student;
  onEdit: (student: Student) => void;
  onDelete: (id: number) => void;
}

const StudentCard: React.FC<StudentCardProps> = ({
  student,
  onEdit,
  onDelete,
}) => {
  const [isDeleting, setIsDeleting] = useState(false);

  const handleDelete = async () => {
    if (window.confirm(`Are you sure you want to delete ${student.name}?`)) {
      setIsDeleting(true);
      onDelete(student.id);
    }
  };

  return (
    <div className="student-card">
      <div className="card-header">
        <h3>{student.name}</h3>
        <span className="id-badge">ID: {student.id}</span>
      </div>
      <div className="card-body">
        <p>
          <strong>Age:</strong> {student.age}
        </p>
        <p>
          <strong>Email:</strong>{" "}
          <span className="email-badge">{student.email}</span>
        </p>
      </div>
      <div className="card-actions">
        <button
          className="edit-btn"
          onClick={() => onEdit(student)}
          title="Edit student"
        >
          ✏️ Edit
        </button>
        <button
          className="delete-btn"
          onClick={handleDelete}
          disabled={isDeleting}
          title="Delete student"
        >
          {isDeleting ? "Deleting..." : "🗑️ Delete"}
        </button>
      </div>
    </div>
  );
};

export default StudentCard;
