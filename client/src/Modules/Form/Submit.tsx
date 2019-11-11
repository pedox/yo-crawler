import { connect, FormikContextType } from 'formik'
import React, { Fragment } from 'react'
import { Button, ButtonProps, Spinner } from 'reactstrap'

interface Props {
  formik: FormikContextType<any>
}

const Submit: React.FC<Props & ButtonProps> = props => {
  const { formik } = props

  const Submitting = () => {
    if (formik.isSubmitting) {
      return (
        <Fragment>
          <Spinner size="sm" />{' '}
        </Fragment>
      )
    }
    return <Fragment></Fragment>
  }

  return (
    <Fragment>
      <Button type="submit" {...props} disabled={formik.isSubmitting}>
        <Submitting />
        {props.children}
      </Button>
    </Fragment>
  )
}

export default connect<ButtonProps>(Submit)
