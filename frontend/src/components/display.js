function linkDisplay({ urls }) {
    return (
        <p>
            {urls.map((url, index) => (
                <span key={index}>
                    <a href={url} target="_blank" rel="noopener noreferrer">{url}</a>
                    {index < urls.length - 1 ? ' -> ' : ''}
                </span>
            ))}
        </p>
    );
}

export default linkDisplay;