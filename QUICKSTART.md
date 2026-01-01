# 🚀 Quick Start Guide

## One-Minute Setup

### Start the Go Backend
```bash
cd c:\Users\Balagowni Jaisheesh\Desktop\Dev\GO\Go_Project
go run cmd/students-api/main.go
```

### Start the React Frontend (in another terminal)
```bash
cd c:\Users\Balagowni Jaisheesh\Desktop\Dev\GO\Go_Project\students-frontend
npm run dev
```

### Open in Browser
Navigate to: **http://localhost:3000**

---

## What You'll See

✅ **Beautiful Student Management Interface**
- Purple gradient background
- Form to add new students (name, age, grade)
- List of students displayed as cards
- Real-time updates when adding students
- Responsive design that works on all devices

## Features to Try

1. **Add a Student**
   - Fill in the form on the left
   - Click "Add Student"
   - See the list update automatically

2. **View Students**
   - Scroll through the student cards
   - See each student's ID, name, age, and grade

3. **Responsive Design**
   - Resize your browser window
   - Watch the layout adapt to mobile, tablet, and desktop

## File Locations

| Item | Location |
|------|----------|
| Frontend Code | `students-frontend/src/` |
| Frontend Docs | `students-frontend/README.md` |
| Backend Code | `cmd/students-api/` |
| Setup Guide | `SETUP.md` |
| Summary | `FRONTEND_SUMMARY.md` |

## Common Commands

### Frontend
```bash
cd students-frontend

# Development with hot reload
npm run dev

# Production build
npm run build

# Preview production build
npm run preview

# Install dependencies
npm install
```

### Backend
```bash
# Run development server
go run cmd/students-api/main.go

# Build executable
go build -o students-api cmd/students-api/main.go

# Run built executable
./students-api
```

## Troubleshooting

### Issue: Frontend shows "Loading..." forever
**Solution**: Make sure the Go backend is running on port 8080

### Issue: Port already in use
**Solution**: 
- Backend: Change port in `config/local.yaml`
- Frontend: Stop other processes using the port

### Issue: Network errors in console
**Solution**: Check that both servers are running and ports are correct

## Tech Stack

- **Backend**: Go with HTTP server and SQLite
- **Frontend**: React 18 + TypeScript + Vite
- **Styling**: CSS3 with gradients and animations
- **API**: REST API with Axios client

## Next Steps

1. ✅ Run both servers successfully
2. ✅ Add some students using the form
3. ✅ Explore the code in `students-frontend/src/`
4. ✅ Try making modifications and see hot reload
5. ✅ Read `FRONTEND_SUMMARY.md` for complete documentation

---

**Enjoy building! 🎉**
