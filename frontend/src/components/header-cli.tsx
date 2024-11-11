'use client';

import React from 'react';

import { IconButton } from '@chakra-ui/react';
import { HiOutlineRefresh } from 'react-icons/hi';

import {requestRest} from '@/services/request.service'

function HeaderCli ({ setOutputValue } : {setOutputValue: React.Dispatch<React.SetStateAction<string> >}) {
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
            onClick={async () => {
              //Limpiar el LocalStorage outputValue, path
              localStorage.setItem('path', '/');
              localStorage.setItem('outputValue', '');
              setOutputValue('');
              //Fetch para reiniciar CLI
              requestRest('GET', '/resetCLI', {})
                .then((data) => {
                  //Si hay errores, mostrarlos en una alerta
                  if (data.errs.length) {
                    alert(
                      data.errs.map((element: string) => {
                        return element + '\n';
                      }),
                    );
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
