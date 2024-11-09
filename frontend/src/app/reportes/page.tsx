'use client'

import React, {useEffect,useState} from 'react';
import { Box,Text } from '@chakra-ui/react';
import ItemContent from '@/components/item-content';

const Reportes = () => {

  type ItemReporte = {
    nombre  :string
    tipo    :string 
    content :string
    key    :string
  }

  const [reportes, setReportes] = useState<ItemReporte[]>([]); 
  const [msgContent, setMsgContent] = useState<string>('No hay Contenido');

  useEffect(() => {
    setMsgContent('Cargando Contenido...');
    //fetchReportes();
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/reportes`, {
      method: 'GET',
    })
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        setReportes(data.files);
        if(data.files.length === 0){
          setMsgContent('No hay Reportes');
        }
      });
  }, []);

  return (
    <>
      <Box
        padding={'2%'}
      >
        {
        reportes.length === 0 ? <Text
          fontSize={'2xl'}
          fontWeight={'bold'}
        >{msgContent}</Text> :
        reportes.map((reporte:ItemReporte) => {
          return (
          <>
            <ItemContent nombre={reporte.nombre} tipo={reporte.tipo} contentRep={reporte.content} key={reporte.key} ></ItemContent>
          </>
        );
        })
        }
      </Box>
    </>
  );
};

export default Reportes;
