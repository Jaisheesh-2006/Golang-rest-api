import React from "react";
import { Student } from "../types";
import "./StudentCard.css";

interface StudentCardProps {
  student: Student;
}

const StudentCard: React.FC<StudentCardProps> = ({ student }) => {
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
    </div>
  );
};

export default StudentCard;
