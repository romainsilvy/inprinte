import React from 'react'
import './App.css';
import Header from './components/Header';
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Footer from './components/Footer';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Header />
        <Routes>
          <Route path="/" element={<h1>Inprinte</h1>} />
          <Route path="/accueil" element={<h1>Accueil</h1>} />
          <Route path="/boutique" element={<h1>Boutique</h1>} />
          <Route path="/propositions" element={<h1>Propositions</h1>} />
          <Route path="/contact" element={<h1>Nous contacter</h1>} />
        </Routes>
      </BrowserRouter>
      <Footer />
    </div>
  );
}

export default App;