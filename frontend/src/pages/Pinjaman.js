import React, { useEffect, useState } from "react";
import api from "../lib/api";

export default function Pinjaman() {
  const [anggota, setAnggota] = useState([]);
  const [buku, setBuku] = useState([]);
  const [overdue, setOverdue] = useState([]);
  const [form, setForm] = useState({ id_anggota: "", id_buku: "" });

  // detail tidak pernah null — selalu object dengan list (aman untuk .map)
  const [detail, setDetail] = useState({ id: null, list: [] });

  useEffect(() => {
    fetchAll();
  }, []);

  function fetchAll() {
    api
      .get("/api/karyawan")
      .then((r) => setAnggota(r.data))
      .catch(() => {});
    api
      .get("/api/buku")
      .then((r) => setBuku(r.data))
      .catch(() => {});
    api
      .get("/api/pinjaman/overdue/anggota")
      .then((r) => setOverdue(r.data))
      .catch(() => {});
  }

  function submit(e) {
    e.preventDefault();
    api
      .post("/api/pinjaman", {
        id_anggota: Number(form.id_anggota),
        id_buku: Number(form.id_buku),
      })
      .then(() => {
        setForm({ id_anggota: "", id_buku: "" });
        fetchAll();
      })
      .catch((err) => alert(err.response?.data?.error || err.message));
  }

  function showDetail(id) {
    api
      .get("/api/pinjaman/anggota/" + id)
      .then((r) =>
        setDetail({
          id,
          list: Array.isArray(r.data) ? r.data : [],
        })
      )
      .catch(() => {
        setDetail({ id: null, list: [] });
      });
  }

  function ret(id) {
    if (!window.confirm("Kembalikan?")) return;
    api
      .put("/api/pinjaman/return/" + id)
      .then(() => {
        // kosongkan detail dan refresh list
        setDetail({ id: null, list: [] });
        fetchAll();
      })
      .catch(() => {
        fetchAll();
        setDetail({ id: null, list: [] });
      });
  }

  function closePopup() {
    setDetail({ id: null, list: [] });
  }

  return (
    <div>
      <h3 className="text-xl font-bold mb-4">Peminjaman</h3>

      {/* Form Pinjam */}
      <form className="flex gap-2 mb-4" onSubmit={submit}>
        <select
          className="border p-2"
          value={form.id_anggota}
          onChange={(e) => setForm({ ...form, id_anggota: e.target.value })}
          required
        >
          <option value="">Pilih Anggota</option>
          {anggota.map((a) => (
            <option value={a.id} key={a.id}>
              {a.nama}
            </option>
          ))}
        </select>

        <select
          className="border p-2"
          value={form.id_buku}
          onChange={(e) => setForm({ ...form, id_buku: e.target.value })}
          required
        >
          <option value="">Pilih Buku</option>
          {buku.map((b) => (
            <option value={b.id} key={b.id}>
              {b.judul}
            </option>
          ))}
        </select>

        <button className="bg-blue-600 text-white px-4">Pinjam</button>
      </form>

      {/* Overdue list */}
      <h4 className="text-lg font-semibold mb-2">
        Peminjam yang Belum Mengembalikan
      </h4>
      <ul className="mb-4">
        {Array.isArray(overdue) &&
          overdue.map((o) => (
            <li
              key={o.id_anggota}
              className="p-2 border mb-1 bg-white flex justify-between"
            >
              <span>
                {o.nama} ({o.divisi})
              </span>
              <button
                className="bg-green-600 text-white px-2"
                onClick={() => showDetail(o.id_anggota)}
              >
                Detail
              </button>
            </li>
          ))}
      </ul>

      {/* Detail popup — hanya ditampilkan kalau ada detail.id */}
      {detail.id && (
        <div className="p-4 border bg-white shadow">
          <h4 className="text-lg font-bold mb-2">Detail Pinjaman</h4>
          <ul>
            {detail.list?.map((p) => (
              <li key={p.id} className="border p-2 mb-1 flex justify-between">
                <div>
                  Buku ID: {p.id_buku} <br />
                  Tgl Pinjam: {new Date(p.tgl_pinjam).toLocaleString()} <br />
                  Status:{" "}
                  {p.status_peminjaman ? "Sudah kembali" : "Belum kembali"}
                </div>
                {!p.status_peminjaman && (
                  <button
                    className="bg-red-600 text-white px-2"
                    onClick={() => ret(p.id)}
                  >
                    Return
                  </button>
                )}
              </li>
            ))}
          </ul>

          <button
            className="bg-gray-700 text-white px-4 mt-2"
            onClick={closePopup}
          >
            Close
          </button>
        </div>
      )}
    </div>
  );
}
