import React, { useState } from 'react';
import { StudentFormData } from '../types';
import { addStudent } from '../services/api';
import './StudentForm.css';

interface StudentFormProps {
  onStudentAdded?: () => void;
}

const StudentForm: React.FC<StudentFormProps> = ({ onStudentAdded }) => {
  const [formData, setFormData] = useState<StudentFormData>({
    name: '',
    age: 0,
    grade: '',
  });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: name === 'age' ? parseInt(value) || 0 : value,
    });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(false);

    try {
      const result = await addStudent(formData);
      if (result) {
        setSuccess(true);
        setFormData({ name: '', age: 0, grade: '' });
        if (onStudentAdded) {
          onStudentAdded();
        }
        setTimeout(() => setSuccess(false), 3000);
      } else {
        setError('Failed to add student');
      }
    } catch (err) {
      setError('Error adding student: ' + (err instanceof Error ? err.message : 'Unknown error'));
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="student-form">
      <div className="form-group">
        <label htmlFor="name">Name:</label>
        <input
          id="name"
          type="text"
          name="name"
          value={formData.name}
          onChange={handleChange}
          placeholder="Enter student name"
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="age">Age:</label>
        <input
          id="age"
          type="number"
          name="age"
          value={formData.age}
          onChange={handleChange}
          placeholder="Enter student age"
          required
          min="1"
        />
      </div>

      <div className="form-group">
        <label htmlFor="grade">Grade:</label>
        <input
          id="grade"
          type="text"
          name="grade"
          value={formData.grade}
          onChange={handleChange}
          placeholder="Enter grade (e.g., A, B, C)"
          required
        />
      </div>

      {error && <div className="error-message">{error}</div>}
      {success && <div className="success-message">Student added successfully!</div>}

      <button type="submit" disabled={loading}>
        {loading ? 'Adding...' : 'Add Student'}
      </button>
    </form>
  );
};

export default StudentForm;
