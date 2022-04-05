import {List, Datagrid, TextField, EditButton, DeleteButton} from 'react-admin';
  
  export const AnalyticsList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="name" />
        <TextField source="is_alive" />
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );