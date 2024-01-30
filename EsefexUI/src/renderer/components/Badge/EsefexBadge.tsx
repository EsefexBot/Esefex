/* eslint-disable jsx-a11y/no-static-element-interactions */
/* eslint-disable jsx-a11y/click-events-have-key-events */
import React from 'react';
import './EsefexBadge.css';

interface EsefexBadgeProps {
  type: 'key-bind' | 'favourite';
  onClick: () => void;
  children: React.Children;
}
function EsefexBadge(props: EsefexBadgeProps) {
  // destructure props if needed

  const { onClick, type, children } = props;

  return (
    <div onClick={onClick} className={`esefex-badge badge-type-${type}`}>
      {children}
    </div>
  );
}

export default EsefexBadge;
