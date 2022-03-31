import React from 'react'
import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Navbar } from './components/Navbar';
// import Footer from './components/Footer'
import Accueil from './pages/Accueil.jsx'
import Boutique from './pages/Boutique'
import Propositions from './pages/Propositions'
import Demande from './pages/Demande'
import Contact from './pages/Contact'
import Login from './pages/Login'
import Panier from './pages/Panier'

function App() {
  return (
    <div className="app">
      <Router>
          <Navbar />
          <Routes>
            <Route path="/" element={<div style= {{display: 'flex', justifyContent: 'center', alignItems: 'center', height: '90vh'}}><h1>Bienvenue chez Inprinte</h1></div>} />
            <Route path="/accueil" element={<Accueil/>} />
            <Route path="/boutique" element={<Boutique/>} />
            <Route path="/propositions" element={<Propositions/>} />
            <Route path="/demande" element={<Demande/>} />
            <Route path="/contact" element={<Contact/>} />
            <Route path="/connexion" element={<Login/>} />
            <Route path="/panier" element={<Panier/>} />
          </Routes>
      </Router>
    </div>
  );
}

export default App;