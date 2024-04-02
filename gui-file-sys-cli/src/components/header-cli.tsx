'use client';

import React from 'react';

import { Button, ChakraProvider } from '@chakra-ui/react';
import { Icon } from '@iconify/react';

const HeaderCli = () => {
  return (
    <div>
      <div className="flex h-[47px] items-center justify-between px-4">
        <div className="hidden md:block">
          <ChakraProvider>
            <Button
              colorScheme='blackAlpha'
              variant={'outline'}
            >
              <Icon icon="codicon:debug-restart" width="24" height="24"  style={{color:'black'}} />
            </Button>
          </ChakraProvider>
        </div>
      </div>
    </div>
  );
};

export default HeaderCli;
