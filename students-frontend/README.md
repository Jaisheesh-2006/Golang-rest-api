# Student Management Frontend

A modern React TypeScript application for managing students, built with Vite and Axios.

## Features

- ✅ **View Students**: Display all students in a beautiful card layout
- ➕ **Add Students**: Form to add new students with name, age, and grade
- 🎨 **Modern UI**: Beautiful gradient design with smooth animations
- 📱 **Responsive Design**: Works perfectly on desktop, tablet, and mobile
- ⚡ **Fast Performance**: Built with Vite for ultra-fast development and production builds
- 🔄 **Real-time Updates**: Student list refreshes automatically after adding new students

## Tech Stack

- **React 18.2** - UI Framework
- **TypeScript** - Type safety
- **Vite** - Build tool
- **Axios** - HTTP client
- **CSS3** - Modern styling with gradients and animations

## Project Structure

```
src/
├── components/          # Reusable React components
│   ├── StudentForm.tsx    # Form for adding new students
│   ├── StudentList.tsx    # List of all students
│   ├── StudentCard.tsx    # Individual student card
│   └── *.css              # Component styles
├── pages/               # Page components
│   ├── Home.tsx          # Main page
│   └── Home.css          # Home page styles
├── services/            # API services
│   └── api.ts            # Axios API client
├── types/               # TypeScript interfaces
│   └── index.ts          # Type definitions
├── App.tsx              # Root component
├── main.tsx             # Entry point
├── index.css            # Global styles
└── App.css              # App styles
```

## Installation

1. **Navigate to the frontend directory:**
   ```bash
   cd students-frontend
   ```

2. **Install dependencies:**
   ```bash
   npm install
   ```

## Development

Start the development server:

```bash
npm run dev
```

The application will be available at `http://localhost:3000`

### Features:
- Hot Module Replacement (HMR) for instant updates
- Proxy to Go backend at `http://localhost:8080`
- TypeScript checking

## Build

Create an optimized production build:

```bash
npm run build
```

The build output will be in the `dist/` directory.

## Preview

Preview the production build locally:

```bash
npm run preview
```

## API Integration

The frontend communicates with the Go backend API:

### Base URL
```
http://localhost:8080/api/students
```

### Endpoints

- **GET** `/api/students` - Get all students
- **GET** `/api/students/{id}` - Get a specific student
- **POST** `/api/students` - Create a new student
- **PATCH** `/api/students/{id}` - Update a student

### Example Request Body (POST/PATCH)
```json
{
  "name": "John Doe",
  "age": 20,
  "grade": "A"
}
```

## Configuration

### Vite Configuration
The Vite configuration includes:
- React plugin setup
- Development server on port 3000
- Proxy to Go backend API

### TypeScript Configuration
- Target: ESNext
- JSX: react-jsx
- Strict mode enabled
- Module resolution: node

## Styling

The application uses:
- **Global styles** in `index.css` with gradient background
- **Component-specific styles** in individual `.css` files
- **Animations** for smooth UI interactions
- **Responsive design** with media queries

### Color Scheme
- Primary Gradient: `#667eea` to `#764ba2` (purple)
- Light backgrounds: `#f5f7fa`
- Dark text: `#333`

## Error Handling

- Network errors are caught and displayed to the user
- User-friendly error messages
- Success notifications for successful operations
- Loading states for async operations

## Browser Support

Works on all modern browsers:
- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## Environment Variables

The frontend requires the Go backend to be running on `http://localhost:8080`.

You can modify the API URL in `src/services/api.ts` if needed.

## Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build locally

## Future Enhancements

- [ ] Edit student functionality
- [ ] Delete student functionality
- [ ] Search and filter students
- [ ] Student details page
- [ ] Export student data
- [ ] Dark mode toggle
- [ ] User authentication

## License

This project is part of the Golang REST API learning project.
