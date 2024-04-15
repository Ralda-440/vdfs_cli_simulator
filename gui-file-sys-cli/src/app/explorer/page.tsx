'use client';

import React, { useEffect, useState } from 'react';

import { Box, IconButton, Input,Alert,AlertIcon,AlertTitle,AlertDescription} from '@chakra-ui/react';
import { SlActionUndo } from 'react-icons/sl';
import ItemContent from '@/components/item-content';

const ExplorerPage = () => {
  const [inputpath, setInputpath] = useState<string>('/');
  const [content, setContent] = useState<Content>({});

  useEffect(() => {
    const path = localStorage.getItem('path');
    if (path) {
      setInputpath(path);
    } else{
      localStorage.setItem('path', '/');
      setInputpath('/');
    }
    fetchExplorer()
  }, []);

  type Content = {
    [key : string]: string;
  }
  
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
    const response = await fetch('http://localhost:4005/explorer', {
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
      <Alert status="error">
        <AlertIcon />
        <AlertTitle mr={2}>Error!</AlertTitle>
        <AlertDescription>
          {data.errs.map((element:Error_) => {
            return element.msg;
          }).join('\n')}
        </AlertDescription>
      </Alert>
      localStorage.setItem('path', '/');
      return;
    }
    setInputpath(path);
    setContent(data.content);
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
        <Box display={'inline-block'} width={'95%'}>
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
      </Box>
      <Box
        padding={'1%'}
      >
        {
          Object.keys(content).map((name) => {
            return (
              <ItemContent key={name} nombre={name} tipo={content[name]} setContent={setContent} fetchExplorer={fetchExplorer} /> 
            );
          })
        }
      </Box>
    </>
  );
};

export default ExplorerPage;
