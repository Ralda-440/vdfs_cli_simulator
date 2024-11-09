'use client';

import React from 'react';

import { IconButton } from '@chakra-ui/react';
import { HiOutlineRefresh } from 'react-icons/hi';

const HeaderCli = () => {
  return (
    <div>
      <div className="flex h-[47px] items-center justify-between">
        <div className="hidden md:block">
          <IconButton
            aria-label="Return"
            icon={
              <HiOutlineRefresh
                style={{
                  color: 'black',
                }}
              />
            }
            variant={'outline'}
            colorScheme="blackAlpha"
            size={'lg'}
            onClick={() => {
              //Limpiar el LocalStorage outputValue, path
              localStorage.setItem('path', '/');
              localStorage.setItem('outputValue', '');
              //Fetch para reiniciar CLI
              fetch(`${process.env.NEXT_PUBLIC_API_URL}/resetCLI`, {
                method: 'GET',
                headers: {
                  'Content-Type': 'application/json',
                },
              })
                .then((res) => {
                  return res.json();
                })
                .then((data) => {
                  //Si hay errores, mostrarlos en una alerta
                  if (data.errs.length) {
                    alert(
                      data.errs.map((element: string) => {
                        return element + '\n';
                      }),
                    );
                  } else {
                    window.location.reload();
                  }
                });
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default HeaderCli;
