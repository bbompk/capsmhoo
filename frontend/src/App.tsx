import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css'
import Navbar from './components/header/Navbar';
import Home from './pages/home/Home';
import CreateProject from './pages/form/createProject';
import CreateTeam from './pages/form/createTeam';
import ProjectList from './pages/proj/ProjectList';
import TeamList from './pages/team/TeamList';
import Login from './pages/login/Login';

function App() {

  return (
    <div>
      <Router>
          <Navbar />

          <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/login' element={<Login />} />
          <Route path='/create-project' element={<CreateProject />} />
          <Route path='/create-team' element={<CreateTeam />} />
          <Route path='/view-project' element={<ProjectList/>} />
          <Route path='/view-team' element={<TeamList/>} />
          
          </Routes>

      </Router>
    </div>
  )
}

export default App
