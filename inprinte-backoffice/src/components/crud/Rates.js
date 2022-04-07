import {List, Datagrid, TextField, Edit, SimpleForm, TextInput, EditButton, Create, DeleteButton, SelectInput} from 'react-admin';
  
  export const RatesList = props => (
    <List {...props}>
      <Datagrid rowClick="edit">
        <TextField source="id" label="Id"/>
        <TextField source="firstname" label="Prénom"/>
        <TextField source="lastname" label="Nom"/>
        <TextField source="id_product" label="Id produit"/>
        <TextField source="product_name" label="Nom du produit"/>
        <TextField source="id_user" label="Id user"/>
        <TextField source="stars_number" label="Note"/>
        <EditButton />
        <DeleteButton />
      </Datagrid>
    </List>
  );

  export const RatesEdit = props => (
    <Edit {...props}>
      <SimpleForm>
      <SelectInput source="stars_number" label="Note" choices={
          [
            { id: '0', name: '0' },
            { id: '1', name: '1' },
            { id: '2', name: '2' },
            { id: '3', name: '3' },
            { id: '4', name: '4' },
            { id: '5', name: '5' },
          ]
        }/>
      </SimpleForm>
    </Edit>
  );

  export const RatesCreate = props => (
    <Create {...props}>
      <SimpleForm>
        <TextInput source="id_user" label="Id user"/>
        <TextInput source="id_product" label="Id produit"/>
        <SelectInput source="stars_number" label="Note" choices={
          [
            { id: '0', name: '0' },
            { id: '1', name: '1' },
            { id: '2', name: '2' },
            { id: '3', name: '3' },
            { id: '4', name: '4' },
            { id: '5', name: '5' },
          ]
        }/>
      </SimpleForm>
    </Create>
  );