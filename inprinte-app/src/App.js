import React from "react";
import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Accueil from "./pages/Accueil.jsx";
import Boutique from "./pages/Boutique";
import Propositions from "./pages/Propositions";
import Demande from "./pages/Demande";
import Contact from "./pages/Contact";
import Panier from "./pages/Panier";
import Connexion from "./pages/Connexion";
import Inscription from "./pages/Inscription";
import Profil from "./pages/Profil";
import Accordion from "./pages/Accordion";

function App() {
  return (
    <div className="app">
      <Router>
        <Routes>
          <Route path="/" element={<Accueil />} />
          <Route path="/boutique" element={<Boutique />} />
          <Route path="/propositions" element={<Propositions />} />
          <Route path="/demande" element={<Demande />} />
          <Route path="/contact" element={<Contact />} />
          <Route path="/faq" element={<Accordion />} />
          <Route path="/profil" element={<Profil />} />
          <Route path="/panier" element={<Panier />} />
          <Route path="/connexion" element={<Connexion />} />
          <Route path="/inscription" element={<Inscription />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
