// components/D3ForceDirectedGraph.jsx

import React, { useRef, useEffect } from 'react';
import * as d3 from 'd3';

const D3ForceDirectedGraph = ({ data, width, height }) => {
  const svgRef = useRef();

  useEffect(() => {
    if (!data || data.length === 0) return;

    const svg = d3.select(svgRef.current);

    // Mengambil data nodes
    const nodes = data.map((d, i) => ({ id: d }));

    // Membuat link berurutan
    const links = [];
    for (let i = 0; i < nodes.length - 1; i++) {
      links.push({ source: nodes[i].id, target: nodes[i + 1].id });
    }

    // Membuat simulasi gaya
    const simulation = d3.forceSimulation(nodes)
      .force('charge', d3.forceManyBody().strength(-100))
      .force('link', d3.forceLink(links).id(d => d.id))
      .force('center', d3.forceCenter(width / 2, height / 2));

    // Menghapus elemen sebelumnya (jika ada)
    svg.selectAll('*').remove();

    // Menggambar link
    const link = svg.selectAll('line')
      .data(links)
      .enter().append('line')
      .attr('stroke', '#999')
      .attr('stroke-opacity', 0.6)
      .attr('stroke-width', 2) // Menetapkan lebar garis
      .attr('x1', d => d.source.x) // Mengatur posisi awal x
      .attr('y1', d => d.source.y) // Mengatur posisi awal y
      .attr('x2', d => d.target.x) // Mengatur posisi akhir x
      .attr('y2', d => d.target.y); // Mengatur posisi akhir y

    // Menggambar node
    const node = svg.selectAll('circle')
      .data(nodes)
      .enter().append('circle')
      .attr('r', 10)
      .attr('fill', 'steelblue')
      .call(d3.drag()
        .on('start', dragstarted)
        .on('drag', dragged)
        .on('end', dragended));

    // Menambahkan label pada node
    node.append('title')
      .text(d => d.id);

    // Menambahkan teks di bawah masing-masing titik
    const text = svg.selectAll('text')
      .data(nodes)
      .enter().append('text')
      .text(d => d.id)
      .attr('x', d => d.x)
      .attr('y', d => d.y + 25)
      .attr('text-anchor', 'middle')
      .attr('fill', 'white')
      .style('font-size', '12px');

    // Memperbarui posisi node dan link setiap iterasi simulasi
    simulation.on('tick', () => {
      link
        .attr('x1', d => d.source.x)
        .attr('y1', d => d.source.y)
        .attr('x2', d => d.target.x)
        .attr('y2', d => d.target.y);

      node
        .attr('cx', d => d.x)
        .attr('cy', d => d.y);

      text
        .attr('x', d => d.x)
        .attr('y', d => d.y + 25);
    });

    // Fungsi untuk memulai drag
    function dragstarted(event, d) {
      if (!event.active) simulation.alphaTarget(0.3).restart();
      d.fx = d.x;
      d.fy = d.y;
    }

    // Fungsi untuk drag
    function dragged(event, d) {
      d.fx = event.x;
      d.fy = event.y;
    }

    // Fungsi untuk mengakhiri drag
    function dragended(event, d) {
      if (!event.active) simulation.alphaTarget(0);
      d.fx = null;
      d.fy = null;
    }

    return () => {
      // Membersihkan simulasi saat komponen dibongkar
      simulation.stop();
    };
  }, [data, width, height]);

  return <svg ref={svgRef} width={width} height={height}></svg>;
};

export default D3ForceDirectedGraph;
