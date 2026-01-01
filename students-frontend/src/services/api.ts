import axios from 'axios';
import { Student, StudentFormData } from '../types';

const API_URL = 'http://localhost:8080/api/students';

export const fetchStudents = async (): Promise<Student[]> => {
  try {
    const response = await axios.get<Student[]>(API_URL);
    return response.data || [];
  } catch (error) {
    console.error('Error fetching students:', error);
    return [];
  }
};

export const addStudent = async (studentData: StudentFormData): Promise<Student | null> => {
  try {
    const response = await axios.post<Student>(API_URL, studentData);
    return response.data;
  } catch (error) {
    console.error('Error adding student:', error);
    return null;
  }
};

export const updateStudent = async (id: number, studentData: Partial<StudentFormData>): Promise<Student | null> => {
  try {
    const response = await axios.patch<Student>(`${API_URL}/${id}`, studentData);
    return response.data;
  } catch (error) {
    console.error('Error updating student:', error);
    return null;
  }
};

export const fetchStudentById = async (id: number): Promise<Student | null> => {
  try {
    const response = await axios.get<Student>(`${API_URL}/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching student:', error);
    return null;
  }
};
