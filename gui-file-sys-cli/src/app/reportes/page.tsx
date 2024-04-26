'use client'

import React, {useEffect,useState} from 'react';
import { Box } from '@chakra-ui/react';
import ItemContent from '@/components/item-content';

const Reportes = () => {

  type ItemReporte = {
    nombre  :string
    tipo    :string 
    content :string
  }

  const [reportes, setReportes] = useState<ItemReporte[]>([]); 

  useEffect(() => {
    //fetchReportes();
    fetch('http://localhost:4005/reportes', {
      method: 'GET',
    })
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        setReportes(data.files);
        console.log(data);
      });
  }, []);

  return (
    <>
      <Box
        padding={'2%'}
      >
        {
        reportes.map((reporte:ItemReporte) => {
          return (
          <>
            <ItemContent nombre={reporte.nombre} tipo={reporte.tipo} contentRep={reporte.content} ></ItemContent>
          </>
        );
        })}
      </Box>
    </>
  );
};

export default Reportes;
