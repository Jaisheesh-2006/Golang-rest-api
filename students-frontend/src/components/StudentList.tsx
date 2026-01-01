import React, { useEffect, useState } from 'react';
import { Student } from '../types';
import StudentCard from './StudentCard';
import { fetchStudents } from '../services/api';
import './StudentList.css';

const StudentList: React.FC = () => {
  const [students, setStudents] = useState<Student[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const getStudents = async () => {
      setLoading(true);
      setError(null);
      try {
        const data = await fetchStudents();
        setStudents(data);
      } catch (err) {
        setError('Failed to fetch students');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    getStudents();
  }, []);

  if (loading) {
    return <div className="loading">Loading students...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (students.length === 0) {
    return <div className="empty-state">No students yet. Add one to get started!</div>;
  }

  return (
    <div className="student-list">
      <h2>Students ({students.length})</h2>
      <div className="cards-container">
        {students.map(student => (
          <StudentCard key={student.id} student={student} />
        ))}
      </div>
    </div>
  );
};

export default StudentList;
