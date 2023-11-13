import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css'
import Navbar from './components/header/Navbar';
import Home from './pages/home/Home';
import CreateTeam from './pages/form/createTeam';
import ProjectList from './pages/proj/ProjectList';
import TeamList from './pages/team/TeamList';
import Register from './pages/register/Register';
import Login from './pages/login/Login';
import Profile from './pages/profile/Profile';
import { AuthProvider } from './components/auth/AuthProvider';
import MockLogin from './components/auth/MockLogin';
import TeamPage from './pages/team/TeamPage';
import MyTeam from './pages/my-team/myTeam';
import ProjectPage from './pages/proj/ProjectPage';
import TeamJoinRequest from './pages/teamJoinRequest/TeamJoinRequest';

function App() {

  return (
    <div>
      <Router>
          <Navbar />

          <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/register' element={<Register/>} />
          <Route path='/login' element={<Login />} />
          <Route path='/profile' element={<AuthProvider><Profile/></AuthProvider>} />
          <Route path='/create-team' element={<AuthProvider><CreateTeam /></AuthProvider>} />
          <Route path='/view-project' element={<AuthProvider><ProjectList/></AuthProvider>} />
          <Route path='/view-team' element={<AuthProvider><TeamList/></AuthProvider>} />
          <Route path='/team-detail/:id' element={<AuthProvider><TeamPage/></AuthProvider>} />
          <Route path='/my-team' element={<AuthProvider><MyTeam/></AuthProvider>} />
          <Route path='/project-detail/:id' element={<AuthProvider><ProjectPage/></AuthProvider>} />
          <Route path='/join-request' element={<AuthProvider><TeamJoinRequest/></AuthProvider>} />
          <Route path='/private' element={<AuthProvider><h1>Private</h1></AuthProvider>} />
          <Route path='/mock-login' element={<MockLogin />} />
          </Routes>

      </Router>
    </div>
  )
}

export default App
