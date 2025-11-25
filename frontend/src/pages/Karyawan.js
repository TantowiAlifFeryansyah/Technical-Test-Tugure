import React, {useEffect, useState} from 'react';
import api from '../lib/api';

export default function Karyawan(){
  const [list,setList]=useState([]);
  const [form,setForm]=useState({nama:'',divisi:''});

  useEffect(()=>{ fetchData() },[]);
  function fetchData(){
    api.get('/api/karyawan').then(r=>setList(r.data));
  }
  function submit(e){
    e.preventDefault();
    api.post('/api/karyawan', form).then(()=>{ 
      setForm({nama:'',divisi:''});
      fetchData();
    });
  }
  function remove(id){
    if(!window.confirm('Hapus?')) return;
    api.delete('/api/karyawan/'+id).then(fetchData);
  }

  return (
    <div>
      <h3 className="text-xl font-bold mb-4">Data Karyawan</h3>

      <form className="flex gap-2 mb-4" onSubmit={submit}>
        <input className="border p-2" placeholder="Nama"
          value={form.nama} onChange={e=>setForm({...form,nama:e.target.value})} required />
        <input className="border p-2" placeholder="Divisi"
          value={form.divisi} onChange={e=>setForm({...form,divisi:e.target.value})} required />
        <button className="bg-blue-600 text-white px-4">Tambah</button>
      </form>

      <table className="w-full border">
        <thead>
          <tr className="bg-gray-200">
            <th className="border p-2">Nama</th>
            <th className="border p-2">Divisi</th>
            <th className="border p-2">Aksi</th>
          </tr>
        </thead>
        <tbody>
          {list.map(i=>(
            <tr key={i.id}>
              <td className="border p-2">{i.nama}</td>
              <td className="border p-2">{i.divisi}</td>
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
