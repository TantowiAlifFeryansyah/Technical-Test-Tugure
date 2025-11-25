import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Karyawan from './pages/Karyawan';
import Buku from './pages/Buku';
import Pinjaman from './pages/Pinjaman';

export default function App(){
  return (
    <BrowserRouter>
      <nav className="bg-blue-600 text-white p-4 flex gap-4">
        <Link to="/" className="hover:underline">Home</Link>
        <Link to="/karyawan" className="hover:underline">Karyawan</Link>
        <Link to="/buku" className="hover:underline">Buku</Link>
        <Link to="/pinjaman" className="hover:underline">Peminjaman</Link>
      </nav>

      <div className="p-6">
        <Routes>
          <Route path="/" element={<h2 className="text-xl font-bold">Tugure Library</h2>} />
          <Route path="/karyawan" element={<Karyawan />} />
          <Route path="/buku" element={<Buku />} />
          <Route path="/pinjaman" element={<Pinjaman />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}
