'use client';

import React, { useEffect, useState } from 'react';

import { Box, IconButton, Input,Button,Text} from '@chakra-ui/react';
import { SlActionUndo } from 'react-icons/sl';
import ItemContent from '@/components/item-content';

const ExplorerPage = () => {
  const [inputpath, setInputpath] = useState<string>('/');
  const [content, setContent] = useState<Map<string,string>>(new Map<string,string>());
  const [isDisabledLogout, setIsDisabledLogout] = useState<boolean>(true);

  useEffect(() => {
    const path = localStorage.getItem('path');
    if (path) {
      setInputpath(path);
    } else{
      localStorage.setItem('path', '/');
      setInputpath('/');
    }
    fetchExplorer()
    //Verificar si hay Login Activo
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/loginActivo`, {
      method: 'GET',
    })
      .then((res) => {
        return res.json();
      })
      .then((data) => {
        if (data.activo) {
          setIsDisabledLogout(false);
        }
      });
  }, []);
  
  type Error_ = {
    msg: string;
    linea: number;
    columna: number;
  };
  

  //Fetch a POST a http://localhost:4005/explorer
  //Body: {path: string} -> obtenido del localStorage
  //Response: {content: map[string]string, errs: []Error_}

  const fetchExplorer = async () => {
    let path = localStorage.getItem('path');
    if (!path) {
      path = '/';
    }
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/explorer`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        path: path,
      }),
    });
    const data = await response.json();
    if (data.errs != null && data.errs.length > 0) {
      alert(data.errs.map((element:Error_) => {
        return element.msg;
      }).join('\n'))
      localStorage.setItem('path', '/');
      return;
    }
    setInputpath(path);
    const JSONData = JSON.stringify(data.content);
    const objectData = JSON.parse(JSONData);
    setContent(new Map(Object.entries(objectData)));
  };

  return (
    <>
      <Box>
        <Box display={'inline-block'} width={'5%'} 
        paddingRight={'1%'}
        >
          <IconButton
            aria-label="Return"
            icon={
              <SlActionUndo
                style={{
                  color: 'black',
                }}
              />
            }
            variant={'outline'}
            colorScheme="blackAlpha"
            size={'lg'}
            onClick={() => {
              let path = localStorage.getItem('path');
              if (path === '/') {
                return;
              }
              path = path?.split('/').slice(0, -1).join('/') || '/';
              localStorage.setItem('path', path);
              setInputpath(path);
              fetchExplorer();
            }}
          />
        </Box>
        <Box display={'inline-block'} width={'75%'}>
          <Input
            width={'100%'}
            size={'lg'}
            focusBorderColor="black"
            isInvalid
            errorBorderColor="gray.400"
            placeholder="/"
            value={inputpath}
            onChange={(e) => setInputpath(e.target.value)}
            readOnly
          ></Input>
        </Box>
        <Box
          display={'inline-block'} width={'20%'}
          textAlign={'right'}
        >
          <Button
            isDisabled={isDisabledLogout}
            colorScheme='red'
            size={'lg'}
            onClick={ async () => {
              setIsDisabledLogout(true);
              const path = localStorage.getItem('path');
              const nameDisk = path?.split('/')[1] || '';
              localStorage.setItem('path', '/'+nameDisk);
              await fetch(`${process.env.NEXT_PUBLIC_API_URL}/logout`, {
                method: 'GET',
              });
              await fetchExplorer();
            }}
          >Logout</Button>
        </Box>
      </Box>
      <Box
        padding={'1%'}
      >
        {
          content.size === 0 ? <Text
            fontSize={'2xl'}
            fontWeight={'bold'}
          >No hay Contenido</Text> :
          Array.from(content).map(([nombre,tipo]) => {
            return (
              <ItemContent key={nombre} nombre={nombre} tipo={tipo} setContent={setContent} fetchExplorer={fetchExplorer} setIsDisabledLogout={setIsDisabledLogout} /> 
          )
          })
        }
      </Box>
    </>
  );
};

export default ExplorerPage;