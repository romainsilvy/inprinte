import {List, Datagrid, TextField, EmailField, UrlField, Edit, SimpleForm, TextInput, Create,} from 'react-admin';
  
  export const UserList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" />
        <TextField source="firstname" />
        <TextField source="lastname" />
        <TextField source="password" />
        <EmailField source="email" />
        <TextField source="address.street" label="Street" />
        <TextField source="address.city" label="City" />
        <TextField source="address.state" label="State" />
        <TextField source="address.country" label="Country" />
        <TextField source="address.zipCode" label="ZipCode" />
        <TextField source="phone" />
      </Datagrid>
    </List>
  );
  
  export const UserEdit = props => (
    <Edit {...props}>
      <SimpleForm>
        <TextInput source="id" disabled />
        <TextInput source="firstname" />
        <TextInput source="lastname" />
        <TextInput source="email" />
        <TextInput source="phone" />
        <TextInput source="password" />
      </SimpleForm>
    </Edit>
  );

  export const UserCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="firstname" />
        <TextInput source="lastname" />
        <TextInput source="email" />
        <TextInput source="phone" />
        <TextInput source="password" />
      </SimpleForm>
    </Create>
  );