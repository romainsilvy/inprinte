import React from 'react'
import './App.css';
import Header from './components/Header';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
// import Footer from './components/Footer'
import {ReactComponent as Logo} from '../../images/inprinte.svg';
import Accueil from './pages/Accueil.jsx'
import Boutique from './pages/Boutique'
import Propositions from './pages/Propositions'
import Demande from './pages/Demande'
import Contact from './pages/Contact'

function App() {
  return (
      <Router>
          <Header />
          <Routes>
            <Route path="/" element={<Logo/>} />
            <Route path="/accueil" element={<Accueil/>} />
            <Route path="/boutique" element={<Boutique/>} />
            <Route path="/propositions" element={<Propositions/>} />
            <Route path="/demande" element={<Demande/>} />
            <Route path="/contact" element={<Contact/>} />
          </Routes>
      </Router>
  );
}

export default App;