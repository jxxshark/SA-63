import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntSpecializedappoint } from '../../api/models/EntSpecializedappoint';
import { EntSpecializeddiag } from '../../api/models/EntSpecializeddiag';
import moment from 'moment';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,

 } from '@backstage/core';
const useStyles = makeStyles({
 table: {
   minWidth: 1900,
 },
});

export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 //const [users, setUsers] = useState(Array);
 const [loading, setLoading] = useState(true);
 const [specialappoint, setSpecialappoint] = useState<EntSpecializedappoint[]>([]);
 const [specializeddiag, setSpecializeddiag] = useState<EntSpecializeddiag[]>([]);
 useEffect(() => {
   const getSpecialappoints = async () => {
     const res = await api.listSpecializedappoint({ limit: 10, offset: 0 });
     setLoading(false);
     setSpecialappoint(res);
   };
   getSpecialappoints();


    const getSpecializeddiag = async () => {
    const res = await api.listSpecializeddiag({ limit: 10, offset: 0 });
    setLoading(false);
    setSpecializeddiag(res);
    console.log(res);
  };
  getSpecializeddiag();
 }, [loading]);

 const deleteSpecialappoints = async (id: number) => {
   const res = await api.deleteSpecializedappoint({ id: id });
   setLoading(true);
 };

 return (
 

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">ผู้ป่วย</TableCell>
           <TableCell align="center">แพทย์ทีทำการนัด</TableCell>
           <TableCell align="center">แพทย์ที่ทำการตรวจ</TableCell>
           <TableCell align="center">วันที่</TableCell>
           <TableCell align="center">ลบข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>

         {specialappoint.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.patient?.name}</TableCell>
             <TableCell align="center">{item.edges?.user?.username}</TableCell>
             {specializeddiag.filter((setfilterid:any) => setfilterid.id === item.edges.specializeddiag.id).map((item2:any) => (
                  <TableCell align="center">{item2.edges.user.username}</TableCell>
              ))}
              <TableCell align="center">{moment(item.deathTime).format('DD/MM/YYYY')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteSpecialappoints(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>

 );

}