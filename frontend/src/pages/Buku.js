import React, {useEffect, useState} from 'react';
import api from '../lib/api';

export default function Buku(){
  const [list,setList]=useState([]);
  const [form,setForm]=useState({judul:'',pengarang:'',kode_buku:'',jumlah_stock:1});

  useEffect(()=>{ fetchData() },[]);
  function fetchData(){
    api.get('/api/buku').then(r=>setList(r.data));
  }
  function submit(e){
    e.preventDefault();
    api.post('/api/buku', {...form, jumlah_stock:Number(form.jumlah_stock)}).then(()=>{
      setForm({judul:'',pengarang:'',kode_buku:'',jumlah_stock:1});
      fetchData();
    });
  }
  function remove(id){
    if(!window.confirm('Hapus?')) return;
    api.delete('/api/buku/'+id).then(fetchData);
  }

  return (
    <div>
      <h3 className="text-xl font-bold mb-4">Data Buku</h3>

      <form className="flex gap-2 mb-4" onSubmit={submit}>
        <input className="border p-2" placeholder="Judul" value={form.judul} onChange={e=>setForm({...form,judul:e.target.value})} required />
        <input className="border p-2" placeholder="Pengarang" value={form.pengarang} onChange={e=>setForm({...form,pengarang:e.target.value})} required />
        <input className="border p-2" placeholder="Kode Buku" value={form.kode_buku} onChange={e=>setForm({...form,kode_buku:e.target.value})} required />
        <input className="border p-2 w-20" type="number" value={form.jumlah_stock} onChange={e=>setForm({...form,jumlah_stock:e.target.value})} required />
        <button className="bg-blue-600 text-white px-4">Tambah</button>
      </form>

      <table className="w-full border">
        <thead>
          <tr className="bg-gray-200">
            <th className="border p-2">Judul</th>
            <th className="border p-2">Kode</th>
            <th className="border p-2">Stock</th>
            <th className="border p-2">Aksi</th>
          </tr>
        </thead>
        <tbody>
          {list.map(i=>(
            <tr key={i.id}>
              <td className="border p-2">{i.judul}</td>
              <td className="border p-2">{i.kode_buku}</td>
              <td className="border p-2">{i.jumlah_stock}</td>
              <td className="border p-2">
                <button className="bg-red-500 text-white px-2" onClick={()=>remove(i.id)}>Hapus</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
