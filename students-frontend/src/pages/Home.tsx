import React, { useState, useCallback } from 'react';
import StudentForm from '../components/StudentForm';
import StudentList from '../components/StudentList';
import './Home.css';

const Home: React.FC = () => {
  const [refreshKey, setRefreshKey] = useState(0);

  const handleStudentAdded = useCallback(() => {
    setRefreshKey(prev => prev + 1);
  }, []);

  return (
    <div className="home">
      <header className="app-header">
        <h1>📚 Student Management System</h1>
        <p>Manage your students efficiently</p>
      </header>

      <div className="container">
        <div className="form-section">
          <StudentForm onStudentAdded={handleStudentAdded} />
        </div>

        <div className="list-section">
          <StudentList key={refreshKey} />
        </div>
      </div>
    </div>
  );
};

export default Home;
