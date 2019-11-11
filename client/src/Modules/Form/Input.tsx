import { connect, FormikContextType, getIn } from 'formik'
import React, { Fragment } from 'react'
import { FormFeedback, Input as BsInput, InputProps } from 'reactstrap'

interface Props {
  name: string
  formik: FormikContextType<any>
}

const Input: React.FC<InputProps & Props> = props => {
  const { name, formik } = props
  const error = getIn(formik.errors, name as string)
  const value = getIn(formik.values, name as string)
  const touch = getIn(formik.touched, name as string)

  return (
    <Fragment>
      <BsInput
        disabled={formik.isSubmitting}
        {...props}
        value={value}
        onChange={formik.handleChange}
        onBlur={formik.handleBlur}
        name={name}
        invalid={touch && error}
      />
      {touch && error ? <FormFeedback>{error}</FormFeedback> : ''}
    </Fragment>
  )
}

export default connect<InputProps & { name: string }>(Input)
