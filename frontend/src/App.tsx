import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css'
import Navbar from './components/header/Navbar';
import Home from './pages/home/Home';
import CreateProject from './pages/form/createProject';
import CreateTeam from './pages/form/createTeam';
import ProjectList from './pages/proj/ProjectList';
import TeamList from './pages/team/TeamList';
import Register from './pages/register/Register';
import Login from './pages/login/Login';
import { AuthProvider } from './components/auth/AuthProvider';
import MockLogin from './components/auth/MockLogin';

function App() {

  return (
    <div>
      <Router>
          <Navbar />

          <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/register' element={<Register/>} />
          <Route path='/login' element={<Login />} />
          {/* <Route path='/create-project' element={<CreateProject />} /> */}
          <Route path='/create-team' element={<CreateTeam />} />
          <Route path='/view-project' element={<AuthProvider><ProjectList/></AuthProvider>} />
          <Route path='/view-team' element={<TeamList/>} />
          <Route path='/private' element={<AuthProvider><h1>Private</h1></AuthProvider>} />
          <Route path='/mock-login' element={<MockLogin />} />
          </Routes>

      </Router>
    </div>
  )
}

export default App
